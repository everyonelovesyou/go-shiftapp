import { setInitialTheme, toggleTheme } from './js/theme.js';

document.addEventListener('DOMContentLoaded', setInitialTheme);

const toggleThemeButton = document.querySelector('#toggle-theme');
toggleThemeButton.addEventListener('click', toggleTheme);