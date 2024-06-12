package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Warehouse(val ID: ULong, val storageID: ULong, val storage: Storage, val workers: List<Worker>?, val manager: Manager)