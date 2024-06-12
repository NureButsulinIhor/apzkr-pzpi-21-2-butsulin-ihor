package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Transfer(val ID: ULong, val carID: ULong, val car: Car, val warehouseID: ULong, val warehouse: Warehouse, val inDate: String, val outDate: String)