package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Slot(val ID: ULong, val maxWeight: Double, val item: Item?, val itemID: ULong?, val device: Device?, val weighingResults: List<WeighingResult>?, val storageID: ULong)