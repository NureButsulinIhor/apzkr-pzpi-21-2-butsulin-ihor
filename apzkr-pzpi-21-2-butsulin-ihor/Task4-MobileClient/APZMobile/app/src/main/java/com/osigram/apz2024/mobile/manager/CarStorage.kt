package com.osigram.apz2024.mobile.manager

import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.pulltorefresh.rememberPullToRefreshState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.lifecycle.viewmodel.compose.viewModel
import com.osigram.apz2024.mobile.auth.LocalAuth
import kotlinx.serialization.Serializable

@Serializable
data class CarStorageRoute(val carID: Long)


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CarStorageScreen(carID: ULong, fromSlot: ULong, onChangeFromSlot: (ULong) -> Unit,
                     toSlot: ULong, onChangeToSlot: (ULong) -> Unit,
                     modifier: Modifier = Modifier, storageViewModel: CarStorageViewModel = viewModel()){
    val token = LocalAuth.current.token
    var isRefreshing by remember{ mutableStateOf(false) }
    var useRefresh by remember{ mutableStateOf(false) }
    val pullToRefreshState = rememberPullToRefreshState()

    LaunchedEffect(true, useRefresh) {
        isRefreshing = true
        storageViewModel.refreshSlots(carID, token)
        isRefreshing = false
    }

    PullToRefreshBox(
        modifier = modifier.fillMaxSize(),
        state = pullToRefreshState,
        isRefreshing = isRefreshing,
        onRefresh = {useRefresh = !useRefresh}
    ){
        LazyColumn(
            modifier = modifier.fillMaxSize()
        ) {
            items(
                items=storageViewModel.slots,
                key={it.ID.toString()}
            ){ slot ->
                SlotElement(slot = slot, fromSlot, onChangeFromSlot,
                    toSlot, onChangeToSlot, modifier = modifier)
            }
        }
    }
}