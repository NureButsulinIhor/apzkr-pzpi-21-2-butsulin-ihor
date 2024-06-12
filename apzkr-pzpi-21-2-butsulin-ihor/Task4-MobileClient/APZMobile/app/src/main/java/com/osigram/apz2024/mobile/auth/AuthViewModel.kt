package com.osigram.apz2024.mobile.auth

import android.content.Context
import android.util.Log
import androidx.compose.ui.platform.LocalContext
import androidx.credentials.CredentialManager
import androidx.credentials.CustomCredential
import androidx.credentials.GetCredentialRequest
import androidx.credentials.GetCredentialResponse
import androidx.credentials.PublicKeyCredential
import androidx.credentials.exceptions.GetCredentialException
import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import com.google.android.libraries.identity.googleid.GetGoogleIdOption
import com.google.android.libraries.identity.googleid.GoogleIdTokenCredential
import com.google.android.libraries.identity.googleid.GoogleIdTokenParsingException
import com.osigram.apz2024.mobile.BuildConfig
import com.osigram.apz2024.mobile.models.Response
import io.ktor.client.HttpClient
import io.ktor.client.call.body
import io.ktor.client.engine.cio.CIO
import io.ktor.client.plugins.contentnegotiation.ContentNegotiation
import io.ktor.client.request.post
import io.ktor.client.request.setBody
import io.ktor.http.ContentType
import io.ktor.http.contentType
import io.ktor.serialization.kotlinx.json.json
import kotlinx.coroutines.launch
import kotlinx.coroutines.runBlocking
import kotlinx.serialization.json.Json

class AuthViewModel: ViewModel(){
    private val client = HttpClient(CIO) {
        install(ContentNegotiation) {
            json(Json{
                ignoreUnknownKeys = true
            })
        }
    }

    suspend fun googleLogin(context: Context): String{
        val credentialManager = CredentialManager.create(context)
        val googleIdOption: GetGoogleIdOption = GetGoogleIdOption.Builder()
            .setFilterByAuthorizedAccounts(false)
            .setServerClientId(BuildConfig.GOOGLE_CLIENT_ID)
//        .setAutoSelectEnabled(true)
//        .setNonce("")
            .build()

        val request: GetCredentialRequest = GetCredentialRequest.Builder()
            .addCredentialOption(googleIdOption)
            .build()

        var backendToken: String
        while (true){
            try {
                val response = credentialManager.getCredential(
                    request = request,
                    context = context,
                )
                backendToken = handleSignIn(response)
                break
            } catch (e: GetCredentialException) {
                Log.e("AuthScreen", e.message?:"")
            } catch (e: AuthException){
                Log.e("AuthScreen", e.message?:"")
            }
        }

        return backendToken
    }

    private suspend fun handleSignIn(result: GetCredentialResponse): String {
        val credential = result.credential
        var responseJson: String = ""

        when (credential) {
            is PublicKeyCredential -> {
                responseJson = credential.authenticationResponseJson
            }

            is CustomCredential -> {
                if (credential.type == GoogleIdTokenCredential.TYPE_GOOGLE_ID_TOKEN_CREDENTIAL) {
                    try {
                        val googleIdTokenCredential = GoogleIdTokenCredential
                            .createFrom(credential.data)
                        responseJson = googleIdTokenCredential.idToken
                    } catch (e: GoogleIdTokenParsingException) {
                        throw AuthException("Received an invalid google id token response")
                    }
                } else {
                    throw AuthException("Unexpected type of credential")
                }
            }

            else -> {
                throw AuthException("Unexpected type of credential")
            }
        }

        return authToBackend(responseJson)
    }

    private suspend fun authToBackend(googleIDToken: String): String{
        val response: Response<ServerResponse> = client.post(BuildConfig.BACKEND_URL + "/login") {
            contentType(ContentType.Application.Json)
            setBody(ServerRequest(googleIDToken))
        }.body()

        if (response.status){
            return response.body?.jwt ?: ""
        } else{
            throw AuthException(response.error)
        }
    }
}