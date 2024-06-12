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
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.launch
import kotlinx.coroutines.withContext
import kotlinx.serialization.Serializable

@Serializable
object TasksRoute


@SuppressLint("UnusedMaterial3ScaffoldPaddingParameter")
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun TasksScreen(fromSlot: ULong, toSlot: ULong, modifier: Modifier = Modifier, tasksViewModel: TasksViewModel = viewModel()){
    val lazyColumnState = rememberLazyListState()
    val expandedFOB by remember{
        derivedStateOf { lazyColumnState.firstVisibleItemIndex == 0 }
    }
    val navController = LocalNavigator.current

    val token = LocalAuth.current.token
    var isRefreshing by remember{ mutableStateOf(false) }
    var useRefresh by remember{ mutableStateOf(false) }
    val pullToRefreshState = rememberPullToRefreshState()

    val coroutineScope = rememberCoroutineScope()
    val onDeleteClick: (ULong) -> Unit = {
        coroutineScope.launch {
            withContext(Dispatchers.IO){
                tasksViewModel.deleteTask(it, token)
            }
        }
    }

    LaunchedEffect(true, useRefresh) {
        isRefreshing = true
        tasksViewModel.refreshTasks(token)
        isRefreshing = false
    }

    Scaffold(
        modifier=modifier.fillMaxSize(),
        floatingActionButton = {
            if (fromSlot != 0UL && toSlot != 0UL){
                ExtendedFloatingActionButton(
                    expanded = expandedFOB,
                    onClick = { navController.navigate(AddTaskRoute) },
                    icon={Icon(Icons.Filled.Add, null)},
                    text = {Text(stringResource(R.string.addTask))}
                )
            }
        }
    ){
        PullToRefreshBox(
            modifier = modifier.fillMaxSize(),
            state = pullToRefreshState,
            isRefreshing = isRefreshing,
            onRefresh = {useRefresh = !useRefresh}
        ){
            LazyColumn(
                modifier = modifier.fillMaxSize(),
                state = lazyColumnState
            ) {
                items(
                    items=tasksViewModel.tasks,
                    key={it.ID.toString()}
                ){
                    TaskElement(it, onDeleteClick, modifier = modifier)
                }
            }
        }
    }

}

@Composable
fun TaskElement(task: Task, onDelete: (ULong) -> Unit, modifier: Modifier = Modifier){
    Column(
        modifier = modifier
    ){
        ListItem(
            overlineContent = {Text(stringResource(R.string.done) + ": " + (if (task.status) stringResource(R.string.yes) else stringResource(R.string.no)))},
            headlineContent = {Text(stringResource(R.string.task) + " #" + task.ID.toString())},
            supportingContent = {Text(stringResource(R.string.from) + ": " + stringResource(R.string.slot).lowercase() + " #" + task.fromSlotID.toString() + "; " +
                    stringResource(R.string.to) + ": " + stringResource(R.string.slot).lowercase() + " #" + task.toSlotID.toString())},
            leadingContent = {
                Icon(
                    Icons.Outlined.Create,
                    null
                )
            },
            trailingContent = {
                IconButton(
                    enabled = !task.status,
                    onClick = { onDelete(task.ID) },

                ) {
                    Icon(Icons.Filled.Delete, contentDescription = null)
                }
            }
        )
        HorizontalDivider()
    }
}