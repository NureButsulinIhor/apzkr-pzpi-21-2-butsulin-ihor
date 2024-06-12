package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Item(val ID: ULong, val name: String, val description: String, val weight: Double, val slotID: ULong?)