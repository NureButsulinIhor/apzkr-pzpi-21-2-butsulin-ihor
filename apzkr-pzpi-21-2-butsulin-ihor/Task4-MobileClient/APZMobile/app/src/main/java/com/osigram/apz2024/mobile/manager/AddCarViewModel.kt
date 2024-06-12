package com.osigram.apz2024.mobile.manager

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import androidx.navigation.NavHostController
import com.osigram.apz2024.mobile.BuildConfig
import com.osigram.apz2024.mobile.auth.AuthException
import com.osigram.apz2024.mobile.auth.ServerRequest
import com.osigram.apz2024.mobile.auth.ServerResponse
import com.osigram.apz2024.mobile.models.Car
import com.osigram.apz2024.mobile.models.Response
import com.osigram.apz2024.mobile.models.Slot
import com.osigram.apz2024.mobile.models.Warehouse
import io.ktor.client.HttpClient
import io.ktor.client.call.NoTransformationFoundException
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.get
import io.ktor.client.request.headers
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.http.ContentType
import io.ktor.http.HttpHeaders
import io.ktor.http.append
import io.ktor.http.contentType
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.asStateFlow
import kotlinx.serialization.Serializable
import kotlinx.serialization.json.Json

@Serializable
data class AddTransferRequest(val outTime: String, val carID: ULong)

class AddCarViewModel: ViewModel() {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json{
                ignoreUnknownKeys = true
            })
        }
    }

    var finished by mutableStateOf(false)
    var _carsState by mutableStateOf(emptyList<Car>())
    val cars: List<Car>
        get() = _carsState

    suspend fun refreshCars(token: String){
        val response: Response<List<Car>> = client.get(BuildConfig.BACKEND_URL + "/admin/car/all") {
            contentType(ContentType.Application.Json)
            headers {
                append(HttpHeaders.Authorization, "Bearer $token")
            }
        }.body()

        if (response.status){
            _carsState = response.body ?: emptyList()
        } else{
            Log.e("refreshCars", "wrong status: " + response.error)
        }
    }

    suspend fun addTransfer(outDate: String, carID: ULong, token: String){
        if (carID == 0UL){
            return
        }

        try{
            val response: Response<Boolean?> = client.post(BuildConfig.BACKEND_URL + "/manager/transfer/") {
                contentType(ContentType.Application.Json)
                headers {
                    append(HttpHeaders.Authorization, "Bearer $token")
                }
                setBody(AddTransferRequest(outDate, carID))
            }.body()

            if (response.status){
                finished = true
            } else{
                Log.e("addTransfer", "wrong status: " + response.error)
            }
        } catch (e: NoTransformationFoundException){
            Log.e("addTransfer", "error: $e")
        }

    }
}