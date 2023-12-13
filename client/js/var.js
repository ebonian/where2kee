let buildings = [];
let toilets = [];
let selectedBuilding = localStorage.getItem("selectedBuilding");
let selectedToilet = null;

const toiletHTMLTemplate = (toilet) => {
  const reviewHTMLTemplate = (review) => `
    <div class="toilet-review-card">
      <div class="toilet-review-content">
        <p class="toilet-review-comment">${review.comment}</p>
        <div class="toilet-card-score">
          <p>${review.score}</p>
          <img
            src="/assets/pile-of-poo-flatten.svg"
            alt="score-icon"
            width="26px"
            height="28px"
          />
        </div>
      </div>
      <div class="toilet-review-user">
        <img
          src="/assets/user-icon.svg"
          alt="user-icon"
          width="22px"
          height="22px"
        />
        <p>${review.username}</p>
      </div>
    </div>
  `;

  const reviewHTML = toilet.reviews.map((review) => reviewHTMLTemplate(review));

  return `
    <div class="toilet-wrapper">
      <div class="toilet-card-container">
        <div class="toilet-card-header">
          <h3 class="toilet-card-title">${toilet.name}</h3>
          <div class="toilet-card-score">
            <p>${toilet.score}</p>
            <img
              src="/assets/pile-of-poo-flatten.svg"
              alt="score-icon"
              width="32px"
            />
          </div>
        </div>
        <p class="toilet-card-description">${toilet.building}</p>
        <div class="toilet-card-image">
          <img src="${toilet.image}" alt="toilet-image" />
        </div>
      </div>
      <div class="toilet-review-container">
        ${reviewHTML.join("")}
      </div>
    </div>
  `;
};
