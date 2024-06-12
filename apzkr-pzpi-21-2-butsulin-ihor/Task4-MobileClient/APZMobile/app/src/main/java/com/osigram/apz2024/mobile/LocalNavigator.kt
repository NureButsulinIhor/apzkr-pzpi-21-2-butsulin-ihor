package com.osigram.apz2024.mobile

import androidx.compose.runtime.compositionLocalOf
import androidx.compose.runtime.staticCompositionLocalOf
import androidx.compose.ui.platform.LocalContext
import androidx.navigation.NavController
import androidx.navigation.NavHostController

var LocalNavigator = staticCompositionLocalOf<NavHostController> { error("No navigator found") }