package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Manager(val ID: ULong, val warehouseID: ULong, val userID: ULong, val user: User)