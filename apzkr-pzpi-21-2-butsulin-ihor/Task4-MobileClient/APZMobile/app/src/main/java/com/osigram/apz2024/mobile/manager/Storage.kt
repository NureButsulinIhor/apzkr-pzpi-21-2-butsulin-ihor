package com.osigram.apz2024.mobile.manager

import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.layout.size
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.automirrored.filled.ArrowBack
import androidx.compose.material.icons.automirrored.filled.ArrowForward
import androidx.compose.material.icons.filled.ArrowForward
import androidx.compose.material.icons.filled.Check
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.FilledTonalButton
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.ListItem
import androidx.compose.material3.Text
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.pulltorefresh.rememberPullToRefreshState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.focusModifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.lifecycle.viewmodel.compose.viewModel
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.auth.LocalAuth
import com.osigram.apz2024.mobile.models.Slot
import kotlinx.serialization.Serializable

@Serializable
object StorageRoute


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun StorageScreen(fromSlot: ULong, onChangeFromSlot: (ULong) -> Unit,
                  toSlot: ULong, onChangeToSlot: (ULong) -> Unit,
                  modifier: Modifier = Modifier, storageViewModel: StorageViewModel = viewModel()
){
    val token = LocalAuth.current.token
    var isRefreshing by remember{ mutableStateOf(false) }
    var useRefresh by remember{ mutableStateOf(false) }
    val pullToRefreshState = rememberPullToRefreshState()

    LaunchedEffect(true, useRefresh) {
        isRefreshing = true
        storageViewModel.refreshSlots(token)
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

@Composable
fun SlotElement(slot: Slot,
                fromSlot: ULong, onChangeFromSlot: (ULong) -> Unit,
                toSlot: ULong, onChangeToSlot: (ULong) -> Unit,
                modifier: Modifier = Modifier){
    val secondaryText: String
    val trailingButton: @Composable () -> Unit
    if (slot.item?.ID != null && slot.item.ID != 0UL){
        secondaryText = stringResource(R.string.item) + ": " + (slot.item.name) + "."
        trailingButton = {
            FilledTonalButton(
                onClick = { if (fromSlot == slot.ID) onChangeFromSlot(0UL) else onChangeFromSlot(slot.ID) },
                contentPadding = ButtonDefaults.ButtonWithIconContentPadding
            ) {
                Icon(
                    if (fromSlot == slot.ID) Icons.Filled.Check else Icons.AutoMirrored.Filled.ArrowForward,
                    contentDescription = null,
                    modifier = modifier.size(ButtonDefaults.IconSize)
                )
                Spacer(modifier.size(ButtonDefaults.IconSpacing))
                Text(stringResource(R.string.from).lowercase())
            }
        }
    } else{
        secondaryText = stringResource(R.string.empty) + "."
        trailingButton = {
            FilledTonalButton(
                onClick = { if (toSlot == slot.ID) onChangeToSlot(0UL) else onChangeToSlot(slot.ID) },
                contentPadding = ButtonDefaults.ButtonWithIconContentPadding
            ) {
                Icon(
                    if (toSlot == slot.ID) Icons.Filled.Check else Icons.AutoMirrored.Filled.ArrowBack,
                    contentDescription = null,
                    modifier = modifier.size(ButtonDefaults.IconSize)
                )
                Spacer(modifier.size(ButtonDefaults.IconSpacing))
                Text(stringResource(R.string.to))
            }
        }
    }

    Column(
        modifier = modifier
    ){
        ListItem(
            headlineContent = {Text(stringResource(R.string.slot) + " #" + slot.ID.toString())},
            supportingContent ={Text(secondaryText)},
            leadingContent = {
                Icon(
                    Icons.Outlined.Place,
                    stringResource(R.string.slot)
                )
            },
            trailingContent = {trailingButton()}
        )
        HorizontalDivider()
    }
}