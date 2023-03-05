import Footer from "./components/Footer";
import Navbar from "./components/Navbar";
import Showcase from "./components/Showcase";

function App() {
  return (
    <>
      <Navbar />
      <main className="h-screen py-24">
        <div className="container px-4 text-center">
          <h1 className="text-slate-900 font-bold text-4xl sm:text-5xl mb-6">I am Cham Manganti, DevOps Engineer</h1>
          <h2 className="text-slate-600 text-lg sm:text-xl max-w-4xl mx-auto">I work on things like CI/CD, kubernetes, on-prem servers, automation, monitoring, and software engineering.</h2>
        </div>
        <div className="container px-4 sm:px-12 md:px-24 lg:px-48 pt-48 pb-24">
          <Showcase />
        </div>
        <Footer />
      </main>
    </>
  );
}

export default App;
