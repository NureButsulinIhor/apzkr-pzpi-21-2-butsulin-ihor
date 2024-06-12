package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class WeighingResult(val ID: ULong, val slotID: ULong, val weight: Double, val time: String)