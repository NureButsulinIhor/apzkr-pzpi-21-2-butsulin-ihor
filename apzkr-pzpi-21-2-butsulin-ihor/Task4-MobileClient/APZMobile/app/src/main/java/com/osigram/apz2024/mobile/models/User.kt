package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class User(val ID: ULong, val email: String, val name: String, val picture: String, val type: String)