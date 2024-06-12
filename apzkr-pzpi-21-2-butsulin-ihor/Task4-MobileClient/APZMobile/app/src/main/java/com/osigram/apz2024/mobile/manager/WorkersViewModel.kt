package com.osigram.apz2024.mobile.manager

import android.util.Log
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.setValue
import androidx.lifecycle.ViewModel
import com.osigram.apz2024.mobile.BuildConfig
import com.osigram.apz2024.mobile.auth.AuthException
import com.osigram.apz2024.mobile.auth.ServerRequest
import com.osigram.apz2024.mobile.auth.ServerResponse
import com.osigram.apz2024.mobile.models.Response
import com.osigram.apz2024.mobile.models.Slot
import com.osigram.apz2024.mobile.models.Warehouse
import com.osigram.apz2024.mobile.models.Worker
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
data class AddTimetableRequest(val workerID: ULong, val startWorkShift: Long, val endWorkShift: Long)

class WorkersViewModel: ViewModel() {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json{
                ignoreUnknownKeys = true
            })
        }
    }

    var _workersState by mutableStateOf(emptyList<Worker>())
    val workers: List<Worker>
        get() = _workersState

    suspend fun refreshWorkers(token: String){
        val response: Response<Warehouse> = client.get(BuildConfig.BACKEND_URL + "/manager/warehouse") {
            contentType(ContentType.Application.Json)
            headers {
                append(HttpHeaders.Authorization, "Bearer $token")
            }
        }.body()

        if (response.status){
            Log.i("refreshWorkers", (response.body?.workers ?: emptyList()).size.toString())
            _workersState = response.body?.workers ?: emptyList()
        } else{
            Log.e("refreshWorkers", "wrong status: " + response.error)
        }
    }

    suspend fun addTimetable(workerID: ULong, token: String){
        try{
            val response: Response<Boolean?> = client.post(BuildConfig.BACKEND_URL + "/manager/timetable/") {
                contentType(ContentType.Application.Json)
                headers {
                    append(HttpHeaders.Authorization, "Bearer $token")
                }
                setBody(AddTimetableRequest(workerID, 0, 16))
            }.body()

            if (response.status){
                refreshWorkers(token)
                return
            } else{
                Log.e("addTimetable", "wrong status: " + response.error)
            }
        } catch (e: NoTransformationFoundException){
            Log.e("addTimetable", "error: $e")
        }

    }
}