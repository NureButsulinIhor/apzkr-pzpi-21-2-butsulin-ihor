package com.osigram.apz2024.mobile.worker

import androidx.appcompat.app.AppCompatDelegate
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.runtime.Composable
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.core.os.LocaleListCompat
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.ui.TopBar

@Composable
fun WorkerApp(modifier: Modifier = Modifier) {
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
        topBar = { TopBar(stringResource(R.string.warehouse), {onChangeLang()}, modifier) }
    ) { innerPadding ->
        Column(modifier = modifier
            .padding(innerPadding)
            .fillMaxSize())
        {
            TasksScreen(modifier)
        }
    }
}