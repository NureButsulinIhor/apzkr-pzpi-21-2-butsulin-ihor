package com.osigram.apz2024.mobile.ui

import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.padding
import androidx.compose.material3.Card
import androidx.compose.material3.Text
import androidx.compose.runtime.Composable
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.text.style.TextAlign
import androidx.compose.ui.unit.dp
import com.osigram.apz2024.mobile.R

@Composable
fun CardWithText(text: String, modifier: Modifier = Modifier){
    Card(
        modifier
            .height(100.dp)
            .padding(2.dp)
    ){
        Box(modifier.fillMaxSize()){
            Text(
                text,
                textAlign = TextAlign.Center,
                modifier = modifier.align(Alignment.Center)
            )
        }
    }
}