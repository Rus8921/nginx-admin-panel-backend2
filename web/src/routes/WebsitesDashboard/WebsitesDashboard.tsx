import "./WebsitesDashboard.css";
import Navigation from "../../components/Navigation/Navigation";
import Dashboard from "./Dashboard/Dashboard";

function WebsitesDashboard() {
  return (
    <div className="WebsitesDashboard">
      <Navigation />
      <Dashboard />
    </div>
  );
}

export default WebsitesDashboard;
