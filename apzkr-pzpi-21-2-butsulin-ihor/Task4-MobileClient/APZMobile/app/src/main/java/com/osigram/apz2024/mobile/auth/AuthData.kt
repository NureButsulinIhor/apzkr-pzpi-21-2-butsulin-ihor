package com.osigram.apz2024.mobile.auth

import androidx.compose.runtime.compositionLocalOf
import androidx.compose.runtime.mutableStateOf
import com.auth0.android.jwt.JWT

class AuthData(){
    var token: String = ""
    var email: String = ""
    var ID: Long = 0
    var name: String = ""
    var picture: String = ""
    var type: String = ""

    constructor(t: String): this(){
        token = t
        val jwt = JWT(token)

        ID = jwt.getClaim("id").asLong() ?: 0
        email = jwt.getClaim("email").asString() ?: ""
        name = jwt.getClaim("name").asString() ?: ""
        picture = jwt.getClaim("picture").asString() ?: ""
        type = jwt.getClaim("type").asString() ?: ""
    }
}

val LocalAuth = compositionLocalOf { AuthData() }