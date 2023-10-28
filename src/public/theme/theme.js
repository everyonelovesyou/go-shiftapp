function setTheme(themeName) {
  localStorage.setItem('theme', themeName);
  document.documentElement.className = themeName;
}

export function toggleTheme() {
  localStorage.getItem('theme') == 'theme-dark'
    ? setTheme('theme-light')
    : setTheme('theme-dark');
}

export function setInitialTheme() {
  window.matchMedia('(prefers-color-scheme: dark)').matches
    ? setTheme('theme-dark')
    : setTheme('theme-light');
}
