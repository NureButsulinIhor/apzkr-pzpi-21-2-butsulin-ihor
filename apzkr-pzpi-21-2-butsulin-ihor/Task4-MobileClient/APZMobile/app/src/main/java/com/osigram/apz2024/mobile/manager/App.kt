package com.osigram.apz2024.mobile.manager

import android.util.Log
import androidx.appcompat.app.AppCompatDelegate
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.outlined.AccountCircle
import androidx.compose.material.icons.outlined.Build
import androidx.compose.material.icons.outlined.Create
import androidx.compose.material.icons.outlined.Place
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.core.os.LocaleListCompat
import androidx.navigation.NavGraph.Companion.findStartDestination
import androidx.navigation.compose.NavHost
import androidx.navigation.compose.composable
import androidx.navigation.toRoute
import com.osigram.apz2024.mobile.LocalNavigator
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.ui.BottomBar
import com.osigram.apz2024.mobile.ui.Route
import com.osigram.apz2024.mobile.ui.TopBar

@Composable
fun ManagerApp(modifier: Modifier = Modifier) {
    val navController = LocalNavigator.current
    val buttons = listOf(
        Route(text=stringResource(R.string.storage), route=StorageRoute, icon=Icons.Outlined.Place),
        Route(text=stringResource(R.string.cars), route=CarsRoute, icon=Icons.Outlined.Build),
        Route(text=stringResource(R.string.tasks), route=TasksRoute, icon=Icons.Outlined.Create),
        Route(text=stringResource(R.string.workers), route=WorkersRoute, icon=Icons.Outlined.AccountCircle)
    )
    val onPageChange: (Route) -> Unit = {
        navController.navigate(it.route){
            // Pop up to the start destination of the graph to
            // avoid building up a large stack of destinations
            // on the back stack as users select items
            popUpTo(navController.graph.findStartDestination().id) {
                saveState = true
            }
            // Avoid multiple copies of the same destination when
            // reselecting the same item
            launchSingleTop = true
            // Restore state when reselecting a previously selected item
            restoreState = true

        }
    }

    var fromSlot by remember {
        mutableStateOf(0UL)
    }
    val onChangeFromSlot: (ULong) -> Unit = {
        fromSlot = it
    }
    var toSlot by remember {
        mutableStateOf(0UL)
    }
    val onChangeToSlot: (ULong) -> Unit = {
        toSlot = it
    }

    val onChangeLang: () -> Unit = {
        var locales = AppCompatDelegate.getApplicationLocales()
        if (locales.toLanguageTags() == "en"){
            locales = LocaleListCompat.forLanguageTags("uk")
        } else{
            locales = LocaleListCompat.forLanguageTags("en")
        }

        AppCompatDelegate.setApplicationLocales(locales)
    }

    Scaffold(
        modifier = modifier.fillMaxSize(),
        bottomBar = { BottomBar(buttonNames = buttons, onChange = onPageChange) },
        topBar = {TopBar(stringResource(R.string.warehouse), {onChangeLang()}, modifier)}
    ) { innerPadding ->
        NavHost(navController = navController, startDestination = StorageRoute, modifier = modifier
            .padding(innerPadding)
            .fillMaxSize()){
            composable<StorageRoute>{
                StorageScreen(fromSlot, onChangeFromSlot, toSlot, onChangeToSlot, modifier)
            }
            composable<CarsRoute> {
                CarsScreen(modifier)
            }
            composable<TasksRoute> {
                TasksScreen(fromSlot, toSlot, modifier)
            }
            composable<AddCarRoute> {
                AddCarScreen(modifier)
            }
            composable<TransfersListRoute> {
                val route: TransfersListRoute = it.toRoute()
                TransfersListScreen(carID = route.carID.toULong(), modifier)
            }
            composable<CarStorageRoute> {
                val route: CarStorageRoute = it.toRoute()
                CarStorageScreen(carID = route.carID.toULong(), fromSlot, onChangeFromSlot, toSlot, onChangeToSlot, modifier)
            }
            composable<AddTaskRoute> {
                AddTaskScreen(fromSlotID = fromSlot, onChangeFromSlot, toSlotID = toSlot, onChangeToSlot, modifier=modifier)
            }
            composable<WorkersRoute> {
                WorkersScreen(modifier)
            }
        }

    }
}

@Preview(showBackground = true)
@Composable
fun ManagerAppPreview() {
    ManagerApp()
}