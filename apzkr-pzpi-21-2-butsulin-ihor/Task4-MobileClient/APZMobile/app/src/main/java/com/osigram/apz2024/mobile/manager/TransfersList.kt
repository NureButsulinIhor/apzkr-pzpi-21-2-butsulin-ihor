package com.osigram.apz2024.mobile.manager

import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.DateRange
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.ExperimentalMaterial3Api
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
import com.osigram.apz2024.mobile.models.Transfer
import kotlinx.serialization.Serializable
import java.time.LocalDate
import java.time.ZonedDateTime
import java.time.format.DateTimeFormatter

@Serializable
data class TransfersListRoute(val carID: Long)


@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TransfersListScreen(carID: ULong, modifier: Modifier = Modifier, transfersViewModel: TransfersListViewModel = viewModel()){
    val token = LocalAuth.current.token
    var isRefreshing by remember{ mutableStateOf(false) }
    var useRefresh by remember{ mutableStateOf(false) }
    val pullToRefreshState = rememberPullToRefreshState()

    LaunchedEffect(true, useRefresh) {
        isRefreshing = true
        transfersViewModel.refreshTransfers(carID, token)
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
                items=transfersViewModel.transfers,
                key={it.ID.toString()}
            ){
                TransferElement(transfer = it, modifier = modifier)
            }
        }
    }
}

@Composable
fun TransferElement(transfer: Transfer, modifier: Modifier = Modifier){
    val inDate = ZonedDateTime.parse(transfer.inDate).toLocalDate().format(DateTimeFormatter.ofPattern("yyyy.MM.dd"))
    val outDate = ZonedDateTime.parse(transfer.outDate).toLocalDate().format(DateTimeFormatter.ofPattern("yyyy.MM.dd"))

    Column(
        modifier = modifier
    ){
        ListItem(
            overlineContent = {Text(stringResource(R.string.warehouse) + " #" + transfer.warehouseID.toString())},
            headlineContent = {Text(stringResource(R.string.transfer) + " #" + transfer.ID.toString())},
            supportingContent ={Text(stringResource(R.string.from) + " " + inDate + " " + stringResource(R.string.to) + " " + outDate)},
            leadingContent = {
                Icon(
                    Icons.Outlined.DateRange,
                    stringResource(R.string.transfer)
                )
            }
        )
        HorizontalDivider()
    }
}