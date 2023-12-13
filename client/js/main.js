const fetchData = async () => {
  buildings = await fetchHandler("/building");
  toilets = await fetchHandler("/toilet");

  if (!selectedBuilding) {
    selectedBuilding = buildings[0].name;
    localStorage.setItem("selectedBuilding", selectedBuilding);
  }

  setBuildingButton(buildings);
  fetchToiletByBuilding(selectedBuilding);
};

fetchData();
