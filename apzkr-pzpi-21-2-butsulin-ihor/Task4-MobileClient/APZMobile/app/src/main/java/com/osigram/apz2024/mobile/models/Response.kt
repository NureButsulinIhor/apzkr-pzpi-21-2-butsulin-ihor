package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Response<T>(val status: Boolean, val error: String = "", val body: T? = null)