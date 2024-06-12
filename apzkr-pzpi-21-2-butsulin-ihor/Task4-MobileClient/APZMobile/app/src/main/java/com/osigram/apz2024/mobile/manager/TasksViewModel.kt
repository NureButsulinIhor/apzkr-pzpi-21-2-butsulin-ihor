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
import com.osigram.apz2024.mobile.models.Task
import com.osigram.apz2024.mobile.models.Warehouse
import io.ktor.client.HttpClient
import io.ktor.client.call.NoTransformationFoundException
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.delete
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
data class TasksResponse(val tasks: List<Task>?)

@Serializable
data class DeleteTaskRequest(val taskID: ULong)

class TasksViewModel: ViewModel() {
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json{
                ignoreUnknownKeys = true
            })
        }
    }

    var _tasksState by mutableStateOf(emptyList<Task>())
    val tasks: List<Task>
        get() = _tasksState

    suspend fun refreshTasks(token: String){
        val response: Response<TasksResponse> = client.get(BuildConfig.BACKEND_URL + "/manager/task/all") {
            contentType(ContentType.Application.Json)
            headers {
                append(HttpHeaders.Authorization, "Bearer $token")
            }
        }.body()

        if (response.status){
            _tasksState = response.body?.tasks ?: emptyList()
        } else{
            Log.e("refreshTasks", "wrong status: " + response.error)
        }
    }

    suspend fun deleteTask(taskID: ULong, token: String){
        if (taskID == 0UL){
            return
        }

        try{
            val response: Response<Boolean?> = client.delete(BuildConfig.BACKEND_URL + "/manager/task/") {
                contentType(ContentType.Application.Json)
                headers {
                    append(HttpHeaders.Authorization, "Bearer $token")
                }
                setBody(DeleteTaskRequest(taskID))
            }.body()

            if (response.status){
                refreshTasks(token)
            } else{
                Log.e("deleteTask", "wrong status: " + response.error)
            }
        } catch (e: NoTransformationFoundException){
            Log.e("deleteTask", "error: $e")
        }
    }
}