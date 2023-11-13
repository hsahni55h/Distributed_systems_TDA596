document.addEventListener('DOMContentLoaded', function () {
    // Display existing VIP names on page load
    displayNames();

    // Form submission handler
    document.getElementById('nameForm').addEventListener('submit', function (event) {
        event.preventDefault();
        addName();
        displayNames();
    });
});

function displayNames() {
    // Fetch VIP names from local storage and update displayBox
    const names = getStoredNames();
    const displayBox = document.getElementById('displayBox');
    displayBox.innerHTML = '<h2>VIP Names:</h2>';
    names.forEach(name => {
        const p = document.createElement('p');
        p.textContent = name;
        displayBox.appendChild(p);
    });
}

function addName() {
    // Add a new VIP name to local storage
    const newName = document.getElementById('newName').value;
    const names = getStoredNames();
    names.push(newName);
    setStoredNames(names);
    document.getElementById('newName').value = ''; // Clear the input field
}

function getStoredNames() {
    // Retrieve names from local storage
    const storedNames = localStorage.getItem('vipNames');
    return storedNames ? JSON.parse(storedNames) : [];
}

function setStoredNames(names) {
    // Store names in local storage
    localStorage.setItem('vipNames', JSON.stringify(names));
}
