package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Storage(val ID: ULong, val slots: List<Slot>?, val type: String)