import { setInitialTheme, toggleTheme } from './theme/theme.js';

document.addEventListener('DOMContentLoaded', setInitialTheme);

const toggleThemeButton = document.querySelector('#toggle-theme');
toggleThemeButton.addEventListener('click', toggleTheme);
