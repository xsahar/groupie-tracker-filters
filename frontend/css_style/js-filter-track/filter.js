// When the window loads, execute the following functions
window.onload = function() {
  slideOne(); // Call the slideOne function
  slideTwo(); // Call the slideTwo function
  slideThree(); // Call the slideThree function
  slideFour(); // Call the slideFour function
}

// Get references to the HTML elements for the sliders and their corresponding display values
let sliderOne = document.getElementById("slider-1");
let sliderTwo = document.getElementById("slider-2");
let sliderThree = document.getElementById("slider-3");
let sliderFour = document.getElementById("slider-4");
let displayValOne = document.getElementById("range1");
let displayValTwo = document.getElementById("range2");
let displayValThree = document.getElementById("range3");
let displayValFour = document.getElementById("range4");

// Set a minimum gap between the sliders
let minGap = 0;

// Get references to the slider tracks
let sliderTrack = document.querySelector(".slider-track");
let sliderTrack2 = document.querySelector(".slider-track2");

// Get the maximum value for the sliders
let sliderMaxValue = document.getElementById("slider-1").max;
let sliderMaxValueTwo = document.getElementById("slider-3").max;

// Function to handle the movement of the first slider
function slideOne() {
  // If the gap between the first and second sliders is less than the minimum gap, adjust the first slider's value
  if (parseInt(sliderTwo.value) - parseInt(sliderOne.value) <= minGap) {
    sliderOne.value = parseInt(sliderTwo.value) - minGap;
  }
  // Update the display value for the first slider
  displayValOne.textContent = sliderOne.value;
  // Call the fillColor function to update the color of the slider track
  fillColor();
}

// Function to handle the movement of the third slider
function slideThree() {
  // If the gap between the third and fourth sliders is less than the minimum gap, adjust the third slider's value
  if (parseInt(sliderFour.value) - parseInt(sliderThree.value) <= minGap) {
    sliderThree.value = parseInt(sliderFour.value) - minGap;
  }
  // Update the display value for the third slider
  displayValThree.textContent = sliderThree.value;
  // Call the fillColor2 function to update the color of the slider track
  fillColor2();
}

// Function to handle the movement of the second slider
function slideTwo() {
  // If the gap between the first and second sliders is less than the minimum gap, adjust the second slider's value
  if (parseInt(sliderTwo.value) - parseInt(sliderOne.value) <= minGap) {
    sliderTwo.value = parseInt(sliderOne.value) + minGap;
  }
  // Update the display value for the second slider
  displayValTwo.textContent = sliderTwo.value;
  // Call the fillColor function to update the color of the slider track
  fillColor();
}

// Function to handle the movement of the fourth slider
function slideFour() {
  // If the gap between the third and fourth sliders is less than the minimum gap, adjust the fourth slider's value
  if (parseInt(sliderFour.value) - parseInt(sliderThree.value) <= minGap) {
    sliderFour.value = parseInt(sliderThree.value) + minGap;
  }
  // Update the display value for the fourth slider
  displayValFour.textContent = sliderFour.value;
  // Call the fillColor2 function to update the color of the slider track
  fillColor2();
}

// Function to update the color of the first and second slider track
function fillColor() {
  // Calculate the percentage of the first and second sliders' values
  percent1 = (sliderOne.value / sliderMaxValue) * 100;
  percent2 = (sliderTwo.value / sliderMaxValue) * 100;
  // Set the background of the first slider track using a linear gradient
  sliderTrack.style.background = `linear-gradient(to right, #dadae5 ${percent1}% , #3264fe ${percent1}% , #3264fe ${percent2}%, #dadae5 ${percent2}%)`;
}

// Function to update the color of the third and fourth slider track
function fillColor2() {
  // Calculate the percentage of the third and fourth sliders' values
  percent3 = (sliderThree.value / sliderMaxValueTwo) * 100;
  percent4 = (sliderFour.value / sliderMaxValueTwo) * 100;
  // Set the background of the second slider track using a linear gradient
  sliderTrack2.style.background = `linear-gradient(to right, #dadae5 ${percent3}% , #3264fe ${percent3}% , #3264fe ${percent4}%, #dadae5 ${percent4}%)`;
}