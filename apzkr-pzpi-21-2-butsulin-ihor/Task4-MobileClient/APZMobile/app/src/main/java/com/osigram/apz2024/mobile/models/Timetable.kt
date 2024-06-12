package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Timetable(val ID: ULong, val workerID: ULong, val startTime: String, val endTime: String)