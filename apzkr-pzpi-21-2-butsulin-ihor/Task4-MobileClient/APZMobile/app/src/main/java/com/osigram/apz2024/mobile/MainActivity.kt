package com.osigram.apz2024.mobile

import android.os.Bundle
import androidx.activity.ComponentActivity
import androidx.activity.compose.setContent
import androidx.activity.enableEdgeToEdge
import androidx.appcompat.app.AppCompatActivity
import androidx.appcompat.app.AppCompatDelegate
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Scaffold
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.runtime.CompositionLocalProvider
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.saveable.rememberSaveable
import androidx.compose.runtime.setValue
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.tooling.preview.Preview
import androidx.core.os.LocaleListCompat
import androidx.navigation.compose.rememberNavController
import com.osigram.apz2024.mobile.auth.AuthData
import com.osigram.apz2024.mobile.auth.AuthScreen
import com.osigram.apz2024.mobile.auth.LocalAuth
import com.osigram.apz2024.mobile.manager.ManagerApp
import com.osigram.apz2024.mobile.ui.CardWithText
import com.osigram.apz2024.mobile.ui.theme.APZMobileTheme
import com.osigram.apz2024.mobile.worker.WorkerApp

class MainActivity : AppCompatActivity() {
    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        enableEdgeToEdge()
        setContent {
            APZMobileTheme {
                App()
            }
        }
    }
}

@Composable
fun App(modifier: Modifier = Modifier){
    var authData by remember {
        mutableStateOf(AuthData())
    }
    val navController = rememberNavController()

    if (authData.ID == 0L){
        Scaffold(modifier = modifier.fillMaxSize()) { innerPadding ->
            Column(
                modifier.padding(innerPadding)
            ) {
                AuthScreen(setAuthData = {authData = it})
            }
        }
    } else{
        CompositionLocalProvider(LocalNavigator provides navController) {
            CompositionLocalProvider(LocalAuth provides authData) {
                if (authData.type == "manager"){
                    ManagerApp(modifier)
                } else if (authData.type == "worker"){
                    WorkerApp(modifier)
                } else {
                    Scaffold(modifier = modifier.fillMaxSize()) { innerPadding ->
                        Column(modifier.padding(innerPadding)){
                            CardWithText(text = stringResource(R.string.errWrongUserType))
                        }
                    }
                }
            }
        }
    }
}

@Preview(showBackground = true)
@Composable
fun AppPreview() {
    APZMobileTheme {
        App()
    }
}