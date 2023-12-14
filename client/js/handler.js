const fetchHandler = async (path) => {
  try {
    const response = await fetch("http://localhost:4000" + path);
    const json = await response.json();
    return json;
  } catch (error) {
    console.error(error);
  }
};

const fetchToiletByBuilding = async (buildingName) => {
  try {
    const buildingId = buildings.find(
      (building) => building.name === buildingName
    ).id;

    const data = await fetchHandler("/toilet/" + buildingId);

    selectedToilet = data;
  } catch (error) {
    console.error(error);
  }
};

const addBuildingButtonEvent = () => {
  const buttons = Array.from(
    document.querySelector(".building-container").children
  );

  buttons.forEach((button) => {
    button.addEventListener("click", () => {
      const buildingName = button.textContent;
      setSelectedBuilding(buildingName);
      mapDataToCard();
    });
  });
};

const setBuildingButton = async (buildings) => {
  const buildingContainer = document.querySelector(".building-container");
  const selectedBuilding = localStorage.getItem("selectedBuilding");

  buildingContainer.innerHTML = "";

  buildings.forEach((building) => {
    const button = document.createElement("button");
    button.textContent = building.name;
    button.className = "building-button";
    if (selectedBuilding === building.name) {
      button.classList.add("building-button-selected");
    }
    buildingContainer.appendChild(button);
  });

  addBuildingButtonEvent();
};

const setSelectedBuilding = (buildingName) => {
  const buttons = Array.from(
    document.querySelector(".building-container").children
  );

  buttons.forEach((button) => {
    if (button.textContent === buildingName) {
      button.classList.add("building-button-selected");
      localStorage.setItem("selectedBuilding", buildingName);
      selectedBuilding = buildingName;
    } else {
      button.classList.remove("building-button-selected");
    }
  });

  fetchToiletByBuilding(buildingName);
};

const mapDataToCard = () => {
  const toiletContainer = document.querySelector(".toilet-container");
  toiletContainer.innerHTML = "";
  selectedToilet.forEach((toilet) => {
    toiletContainer.innerHTML += toiletHTMLTemplate(toilet);
  });
}
