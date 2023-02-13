import React from "react";

function Dashboard() {
  return (
    <div>
      <div>
        <h1>Allergy summary</h1>
        <div>
          <ul>
            <li>Allergy Recorded : 55</li>
            <li>Allergy Cured: 20</li>
            <li>Allergy Left: 10</li>
          </ul>

          <div>Bar Graph Here</div>
        </div>
      </div>
      <div>
        <h1>Allergies</h1>
        <div>
          <ul>
            <li>
              <div>
                <h2>Head Ache</h2>
                <div>Suggestions : 4/5</div>
              </div>
              <div>Severity Indicator</div>
            </li>
          </ul>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
