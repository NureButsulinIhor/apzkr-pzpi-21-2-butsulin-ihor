package com.osigram.apz2024.mobile.manager

import android.annotation.SuppressLint
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Add
import androidx.compose.material.icons.filled.Delete
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Add
import androidx.compose.material.icons.outlined.Create
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ExtendedFloatingActionButton
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.ListItem
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.material3.pulltorefresh.PullToRefreshBox
import androidx.compose.material3.pulltorefresh.rememberPullToRefreshState
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.derivedStateOf
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.rememberCoroutineScope
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.lifecycle.viewmodel.compose.viewModel
import com.osigram.apz2024.mobile.LocalNavigator
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.auth.LocalAuth
import com.osigram.apz2024.mobile.models.Slot
import com.osigram.apz2024.mobile.models.Task
import com.osigram.apz2024.mobile.models.Worker
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import kotlinx.serialization.Serializable
import java.time.ZonedDateTime

@Serializable
object WorkersRoute


@SuppressLint("UnusedMaterial3ScaffoldPaddingParameter")
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun WorkersScreen(modifier: Modifier = Modifier, workersViewModel: WorkersViewModel = viewModel()){
    val token = LocalAuth.current.token
    var isRefreshing by remember{ mutableStateOf(false) }
    var useRefresh by remember{ mutableStateOf(false) }
    val pullToRefreshState = rememberPullToRefreshState()

    val coroutineScope = rememberCoroutineScope()
    val onAddTimetableClick: (ULong) -> Unit = {
        coroutineScope.launch {
            withContext(Dispatchers.IO){
                workersViewModel.addTimetable(it, token)
            }
        }
    }

    LaunchedEffect(true, useRefresh) {
        isRefreshing = true
        workersViewModel.refreshWorkers(token)
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
                items=workersViewModel.workers,
                key={it.ID.toString()}
            ){
                WorkerElement(it, onAddTimetableClick, modifier = modifier)
            }
        }
    }

}

@Composable
fun WorkerElement(worker: Worker, onAddTimetable: (ULong) -> Unit, modifier: Modifier = Modifier){
    val isWorking = worker.timetables?.filter{
        ZonedDateTime.parse(it.startTime) < ZonedDateTime.now() &&
                ZonedDateTime.parse(it.endTime) > ZonedDateTime.now()
    }?.any() ?: false

    Column(
        modifier = modifier
    ){
        ListItem(
            headlineContent = {Text(stringResource(R.string.worker) + " " + worker.user.email)},
            supportingContent = {Text(stringResource(R.string.status) + ": " + (if (isWorking) stringResource(R.string.working) else stringResource(R.string.notWorking)))},
            leadingContent = {
                Icon(
                    Icons.Outlined.AccountCircle,
                    null
                )
            },
            trailingContent = {
                IconButton(
                    enabled = !isWorking,
                    onClick = { onAddTimetable(worker.ID) },

                ) {
                    Icon(Icons.Filled.Add, contentDescription = null)
                }
            }
        )
        HorizontalDivider()
    }
}