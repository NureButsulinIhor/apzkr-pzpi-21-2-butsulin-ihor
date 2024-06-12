package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Car(val ID: ULong, val storageID: ULong, val storage: Storage, val ownerID: ULong, val owner: User)