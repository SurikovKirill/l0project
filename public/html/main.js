getOrder = async (event) => {
    console.log(event.target[0].value)
    event.preventDefault();
    await axios.get(`http://localhost:8080/order/${event.target[0].value}`).then(response => alert(JSON.stringify(response.data)));
}

document.getElementById('inputForm').addEventListener('submit', getOrder);
