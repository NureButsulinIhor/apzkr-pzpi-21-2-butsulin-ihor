package com.osigram.apz2024.mobile.ui

import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Row
import androidx.compose.material.icons.Icons
import androidx.compose.material.icons.filled.Info
import androidx.compose.material3.Icon
import androidx.compose.material3.NavigationBar
import androidx.compose.material3.NavigationBarItem
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.navigation.NavDestination.Companion.hierarchy
import androidx.navigation.compose.currentBackStackEntryAsState
import com.osigram.apz2024.mobile.LocalNavigator

@Composable
fun BottomBar(buttonNames: List<Route>, onChange: (Route) -> Unit){
    val navController = LocalNavigator.current

    NavigationBar {
        Row(
            horizontalArrangement = Arrangement.Center,
        ) {
            buttonNames.forEach { route ->
                NavigationBarItem(
                    selected = false,
                    onClick = { onChange(route) },
                    label = { Text(route.text) },
                    icon = {
                        Icon(
                            route.icon,
                            contentDescription = route.text,
                        )
                    },
                )
            }
        }
    }
}