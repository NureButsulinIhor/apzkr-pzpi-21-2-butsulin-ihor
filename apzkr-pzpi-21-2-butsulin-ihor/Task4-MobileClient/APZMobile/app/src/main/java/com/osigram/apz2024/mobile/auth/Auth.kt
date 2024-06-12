package com.osigram.apz2024.mobile.auth

import androidx.compose.foundation.layout.Column
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.Modifier
import androidx.compose.ui.platform.LocalContext
import androidx.compose.ui.tooling.preview.Preview
import androidx.lifecycle.viewmodel.compose.viewModel
import com.osigram.apz2024.mobile.R
import com.osigram.apz2024.mobile.ui.CardWithText

@Composable
fun AuthScreen(setAuthData: (AuthData) -> Unit, modifier: Modifier = Modifier, authViewModel: AuthViewModel = viewModel()){
    val context = LocalContext.current

    LaunchedEffect(true){
        setAuthData(AuthData(authViewModel.googleLogin(context)))
    }

    Column(
        modifier
    )
    {
        CardWithText(text = stringResource(R.string.loginWelcome), modifier = modifier)
    }
}

@Preview(showBackground = true)
@Composable
fun AuthScreenPreview() {
    AuthScreen(setAuthData = {})
}


