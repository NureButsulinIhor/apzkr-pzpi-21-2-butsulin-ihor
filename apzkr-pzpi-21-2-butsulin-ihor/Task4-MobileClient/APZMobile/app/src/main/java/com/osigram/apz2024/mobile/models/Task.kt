package com.osigram.apz2024.mobile.models

import kotlinx.serialization.Serializable

@Serializable
data class Task(val ID: ULong, val workerID: ULong, val worker: Worker, val fromSlotID: ULong, val fromSlot: Slot, val toSlotID: ULong, val toSlot: Slot, val status: Boolean)