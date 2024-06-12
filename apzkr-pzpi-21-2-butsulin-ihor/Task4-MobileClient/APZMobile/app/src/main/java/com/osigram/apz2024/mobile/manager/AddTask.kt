package com.osigram.apz2024.mobile.manager

import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Add
import androidx.compose.material.icons.filled.Check
import androidx.compose.material3.DatePicker
import androidx.compose.material3.DisplayMode
import androidx.compose.material3.DropdownMenuItem
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ExposedDropdownMenuBox
import androidx.compose.material3.ExposedDropdownMenuDefaults
import androidx.compose.material3.ExtendedFloatingActionButton
import androidx.compose.material3.Icon
import androidx.compose.material3.MenuAnchorType
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.TextField
import androidx.compose.material3.rememberDatePickerState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.derivedStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.lifecycle.viewmodel.compose.viewModel
import com.osigram.apz2024.mobile.LocalNavigator
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.auth.LocalAuth
import com.osigram.apz2024.mobile.models.Car
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import kotlinx.serialization.Serializable
import java.text.DateFormat
import java.time.Instant
import java.time.format.DateTimeFormatter
import java.util.Calendar
import java.util.Date
import java.util.TimeZone

@Serializable
object AddTaskRoute


@Composable
fun AddTaskScreen(fromSlotID: ULong, onChangeFromSlotID: (ULong) -> Unit, toSlotID: ULong, onChangeToSlotID: (ULong) -> Unit, modifier: Modifier = Modifier, addTaskViewModel: AddTaskViewModel = viewModel()){
    val navController = LocalNavigator.current
    val token = LocalAuth.current.token
    val coroutineScope = rememberCoroutineScope()

    val onCreateClick = {
        coroutineScope.launch {
            withContext(Dispatchers.IO){
                addTaskViewModel.addTask(fromSlotID, onChangeFromSlotID, toSlotID, onChangeToSlotID, token)
            }
        }
    }

    LaunchedEffect(true) {
        addTaskViewModel.finished = false
        addTaskViewModel.refreshSlots(fromSlotID, toSlotID, token)
        if (fromSlotID == 0UL || toSlotID == 0UL){
            addTaskViewModel.finished = true
        }
    }
    LaunchedEffect(addTaskViewModel.finished) {
        if (addTaskViewModel.finished){
            navController.popBackStack()
        }
    }

    Scaffold(
        modifier=modifier.fillMaxSize(),
        floatingActionButton = {
            ExtendedFloatingActionButton(
                onClick = { onCreateClick() },
                icon={ Icon(Icons.Filled.Check, null) },
                text = { Text(stringResource(R.string.done)) }
            )
        }
    ){
        Column (modifier= modifier.padding(it).fillMaxWidth()){
            if (addTaskViewModel._fromSlotState != null){
                SlotElement(
                    slot = addTaskViewModel._fromSlotState!!,
                    fromSlot = fromSlotID,
                    onChangeFromSlot = {},
                    toSlot = toSlotID,
                    onChangeToSlot = {}
                )
            }
            if (addTaskViewModel._toSlotState != null){
                SlotElement(
                    slot = addTaskViewModel._toSlotState!!,
                    fromSlot = fromSlotID,
                    onChangeFromSlot = {},
                    toSlot = toSlotID,
                    onChangeToSlot = {}
                )
            }
        }
    }
}