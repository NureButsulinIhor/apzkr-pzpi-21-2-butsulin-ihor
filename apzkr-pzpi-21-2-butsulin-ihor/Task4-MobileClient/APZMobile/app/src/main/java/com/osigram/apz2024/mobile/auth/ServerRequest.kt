package com.osigram.apz2024.mobile.auth

import kotlinx.serialization.Serializable

@Serializable
internal data class ServerRequest (val googleJWT: String)