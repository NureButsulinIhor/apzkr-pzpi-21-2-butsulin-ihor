package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Worker(val ID: ULong, val userID: ULong, val user: User, val warehouseID: ULong, val timetables: List<Timetable>?)