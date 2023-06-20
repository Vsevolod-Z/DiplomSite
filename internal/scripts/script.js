const cardList = document.querySelector('.card-list');

while(true){
// Создаем новую карточку
const newCard = document.createElement('div');
newCard.classList.add('card-container');
newCard.innerHTML = `
<div class="card-container">
        <a href="/" class="hero-image-container">
        <img class="hero-image" src="https://cdn.akamai.steamstatic.com/steam/apps/730/header.jpg?t=1668125812" alt="Spinning glass cube"/>
        </a>
        <main class="main-content">
            <h1 id="card-name">Counter-Strike: Global Offensive</h1>
        </main>    
</div>
`;

// Добавляем новую карточку в конец списка
cardList.appendChild(newCard);
}