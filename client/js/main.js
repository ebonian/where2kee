const fetchData = async () => {
  buildings = await fetchHandler("/building");
  toilets = await fetchHandler("/toilet");

  if (!selectedBuilding) {
    selectedBuilding = buildings[0].name;
    localStorage.setItem("selectedBuilding", selectedBuilding);
  }
  setLoadingScreen(false);

  await setBuildingButton(buildings);
  await fetchToiletByBuilding(selectedBuilding);

  mapDataToCard();
};

fetchData();
