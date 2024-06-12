package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Device(val ID: String, val slotID: ULong)