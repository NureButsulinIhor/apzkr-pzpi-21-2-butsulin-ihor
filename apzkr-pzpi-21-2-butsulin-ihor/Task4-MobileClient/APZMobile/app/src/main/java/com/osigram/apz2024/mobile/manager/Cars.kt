package com.osigram.apz2024.mobile.manager

import android.annotation.SuppressLint
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.rememberLazyListState
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Add
import androidx.compose.material.icons.filled.Info
import androidx.compose.material.icons.outlined.Build
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.ButtonDefaults
import androidx.compose.material3.ExperimentalMaterial3Api
import androidx.compose.material3.ExtendedFloatingActionButton
import androidx.compose.material3.FilledTonalButton
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
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
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.focus.focusModifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.lifecycle.viewmodel.compose.viewModel
import com.osigram.apz2024.mobile.LocalNavigator
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.auth.LocalAuth
import com.osigram.apz2024.mobile.models.Car
import com.osigram.apz2024.mobile.models.Slot
import kotlinx.serialization.Serializable

@Serializable
object CarsRoute


@SuppressLint("UnusedMaterial3ScaffoldPaddingParameter")
@OptIn(ExperimentalMaterial3Api::class)
@Composable
fun CarsScreen(modifier: Modifier = Modifier, carsViewModel: CarsViewModel = viewModel()){
    val lazyColumnState = rememberLazyListState()
    val expandedFOB by remember{
        derivedStateOf { lazyColumnState.firstVisibleItemIndex == 0 }
    }
    val navController = LocalNavigator.current

    val token = LocalAuth.current.token
    var isRefreshing by remember{ mutableStateOf(false) }
    var useRefresh by remember{ mutableStateOf(false) }
    val pullToRefreshState = rememberPullToRefreshState()

    LaunchedEffect(true, useRefresh) {
        isRefreshing = true
        carsViewModel.refreshCars(token)
        isRefreshing = false
    }

    Scaffold(
        modifier=modifier.fillMaxSize(),
        floatingActionButton = {
            ExtendedFloatingActionButton(
                expanded = expandedFOB,
                onClick = { navController.navigate(AddCarRoute) },
                icon={Icon(Icons.Filled.Add, null)},
                text = {Text(stringResource(R.string.addCar))}
            )
        }
    ) {
        PullToRefreshBox(
            modifier = modifier
//                .padding(it)
                .fillMaxSize(),
            state = pullToRefreshState,
            isRefreshing = isRefreshing,
            onRefresh = {useRefresh = !useRefresh}
        ){
            LazyColumn(
                modifier = modifier.fillMaxSize(),
                state=lazyColumnState
            ) {
                items(
                    items=carsViewModel.cars,
                    key={it.ID.toString()}
                ){ car ->
                    CarElement(car = car, modifier = modifier)
                }
            }
        }
    }

}

@Composable
fun CarElement(car: Car, modifier: Modifier = Modifier){
    val navController = LocalNavigator.current

    Column(
        modifier = modifier
    ){
        ListItem(
            headlineContent = {Text(stringResource(R.string.car) + " #" + car.ID.toString())},
            supportingContent ={Text(stringResource(R.string.owner) + ": " + car.owner.email)},
            leadingContent = {
                Icon(
                    Icons.Outlined.Build,
                    stringResource(R.string.car)
                )
            },
            trailingContent = {
                Column(
                    modifier
                ){
                    FilledTonalButton(
                        onClick = { navController.navigate(TransfersListRoute(car.ID.toLong())) },
                        contentPadding = ButtonDefaults.ButtonWithIconContentPadding,
                        modifier = modifier.align(Alignment.CenterHorizontally)
                    ) {
                        Icon(
                            Icons.Filled.Info,
                            contentDescription = null,
                            modifier = modifier.size(ButtonDefaults.IconSize)
                        )
                        Spacer(modifier.size(ButtonDefaults.IconSpacing))
                        Text(stringResource(R.string.transfers))
                    }
                    FilledTonalButton(
                        onClick = { navController.navigate(CarStorageRoute(car.ID.toLong())) },
                        contentPadding = ButtonDefaults.ButtonWithIconContentPadding,
                        modifier = modifier.align(Alignment.CenterHorizontally)
                    ) {
                        Icon(
                            Icons.Filled.Info,
                            contentDescription = null,
                            modifier = modifier.size(ButtonDefaults.IconSize)
                        )
                        Spacer(modifier.size(ButtonDefaults.IconSpacing))
                        Text(stringResource(R.string.storage))
                    }
                }
            }
        )
        HorizontalDivider()
    }
}

@Preview(showBackground = true)
@Composable
fun CarsScreenPreview() {
    CarsScreen()
}