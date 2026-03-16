// Application Scripts

// Theme Management
const themeToggleBtn = document.getElementById('theme-toggle');
const darkIcon = document.getElementById('theme-toggle-dark-icon');
const lightIcon = document.getElementById('theme-toggle-light-icon');

function updateThemeIcons(theme) {
    if (theme === 'dark') {
        darkIcon.style.display = 'inline';
        lightIcon.style.display = 'none';
    } else {
        darkIcon.style.display = 'none';
        lightIcon.style.display = 'inline';
    }
}

function initTheme() {
    const savedTheme = localStorage.getItem('theme') || (window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light');
    document.documentElement.setAttribute('data-theme', savedTheme);
    updateThemeIcons(savedTheme);
}

if (themeToggleBtn) {
    themeToggleBtn.addEventListener('click', () => {
        const currentTheme = document.documentElement.getAttribute('data-theme');
        const newTheme = currentTheme === 'dark' ? 'light' : 'dark';

        document.documentElement.setAttribute('data-theme', newTheme);
        localStorage.setItem('theme', newTheme);
        updateThemeIcons(newTheme);
    });
}

// Initialize on load
initTheme();

// HTMX Event Listeners
document.addEventListener('htmx:afterRequest', function (evt) {
    if (evt.detail.target.id === 'sync-status') {
        if (evt.detail.successful) {
            console.log("Sync complete. Reloading in 1.5 seconds...");
            setTimeout(() => {
                location.reload();
            }, 1500);
        }
    }
});
