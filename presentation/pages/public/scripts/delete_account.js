let userID;

document.getElementById('loginForm').addEventListener('submit', async function (e) {
    e.preventDefault();

    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;

    const response = await fetch(`/api/user/login`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ email, password })
    });

    if (response.ok) {
        const data = await response.json();
        userID = data["data"]["id"];
        document.getElementById('loginForm').style.display = 'none';
        document.getElementById('deleteSection').style.display = 'block';
        document.getElementById('logoutButton').style.display = 'inline-block';
    } else {
        const data = await response.json();
        console.log(data)
        alert('Login failed\n' + data["message"]);
    }
});

document.getElementById('deleteButton').addEventListener('click', async function () {
    if (confirm("Are you sure you want to delete your account? This action cannot be undone.")) {
        const response = await fetch(`/api/user/${userID}`, {
            method: 'DELETE'
        });

        if (response.ok) {
            alert('Account deleted successfully');
            document.getElementById('deleteSection').style.display = 'none';
            document.getElementById('loginForm').style.display = 'block';
            window.location.reload(); // Reload the page to reset the form
        } else {
            alert('Failed to delete account');
        }
    }
});

document.getElementById('logoutButton').addEventListener('click', async function () {
    // Optionally handle logout logic here, e.g., calling a logout endpoint
    // For simplicity, just reload the page
    window.location.reload();
});
