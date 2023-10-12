
var dropdown = document.getElementsByClassName("dropdown-btn");
var i;

for (i = 0; i < dropdown.length; i++) {
  dropdown[i].addEventListener("click", function() {
    this.classList.toggle("active");
    var dropdownContent = this.nextElementSibling;
    if (dropdownContent.style.display === "block") {
      dropdownContent.style.display = "none";
    } else {
      dropdownContent.style.display = "block";
    }
  });
}

const sidebar = document.querySelector('.sidebar');
const openSidebarButton = document.querySelector('.open-sidebar-button');
const closeSidebarButton = document.querySelector('.close-sidebar-button');

openSidebarButton.addEventListener('click', () => {
    sidebar.style.left = '0';
});

closeSidebarButton.addEventListener('click', () => {
    sidebar.style.left = '-250px';
});
