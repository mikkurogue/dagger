document.addEventListener('DOMContentLoaded', () => {
    const commandElement = document.getElementById('command');
    const commandText = 'dagger';
    const typingSpeed = 150; // Typing speed in milliseconds
    const deleteSpeed = 100; // Deleting speed in milliseconds
    const pauseTime = 4000; // Pause time in milliseconds before deleting

    let isDeleting = false;
    let charIndex = 0;

    function type() {
        if (isDeleting) {
            if (charIndex > 0) {
                charIndex--;
                commandElement.innerHTML = commandText.substring(0, charIndex);
                setTimeout(type, deleteSpeed);
            } else {
                isDeleting = false;
                setTimeout(type, typingSpeed);
            }
        } else {
            if (charIndex < commandText.length) {
                commandElement.innerHTML = commandText.substring(0, charIndex + 1);
                charIndex++;
                setTimeout(type, typingSpeed);
            } else {
                isDeleting = true;
                setTimeout(type, pauseTime);
            }
        }
    }

    type();
});