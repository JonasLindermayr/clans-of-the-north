import { useEffect, useState } from "react";

export function MapView() {
  const [tiles, setTiles] = useState([]);
  const [centerX, setCenterX] = useState(0); // Aktuelle Kartenposition X
  const [centerY, setCenterY] = useState(0); // Aktuelle Kartenposition Y

  useEffect(() => {
    fetchMapData();
  }, [centerX, centerY]);

  const fetchMapData = () => {
    fetch(`http://localhost:8080/map?centerX=${centerX}&centerY=${centerY}`)
      .then((res) => res.json())
      .then((data) => setTiles(data));
  };

  const handleArrowClick = (direction) => {
    switch (direction) {
      case "up":
        setCenterY(centerY - 1);
        break;
      case "down":
        setCenterY(centerY + 1);
        break;
      case "left":
        setCenterX(centerX - 1);
        break;
      case "right":
        setCenterX(centerX + 1);
        break;
      default:
        break;
    }
  };

  return (
    <div>
      <h2>Map</h2>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <div>
          <button onClick={() => handleArrowClick("up")}>↑</button>
        </div>
        <div style={{ display: "flex" }}>
          <button onClick={() => handleArrowClick("left")}>←</button>
          <div
            style={{
              fontFamily: "monospace",
              whiteSpace: "pre",
              fontSize: "30px",
            }}
          >
            {renderMap(tiles, centerX, centerY)}{" "}
          </div>
          <button onClick={() => handleArrowClick("right")}>→</button>
        </div>
        <div>
          <button onClick={() => handleArrowClick("down")}>↓</button>
        </div>
      </div>
    </div>
  );
}

function renderMap(tiles, centerX, centerY) {
  const map = {};
  const tileArray = Object.values(tiles);

  tileArray.forEach((tile) => {
    const key = `${tile.X},${tile.Y}`;
    let symbol = ".";

    if (tile.villageID) symbol = "🏰";
    else if (tile.Terrain === "forest") symbol = "🌲";
    else if (tile.Terrain === "mountain") symbol = "⛰️";
    else if (tile.Terrain === "lake") symbol = "💧";
    else if (tile.Terrain === "village_spot") symbol = "⚪";
    else symbol = "🟩";

    map[key] = symbol;
  });

  let output = "";
  for (let y = -5; y <= 5; y++) {
    for (let x = -5; x <= 5; x++) {
      output += map[`${x + centerX},${y + centerY}`] || " ";
    }
    output += "\n";
  }

  return output;
}
