# Ascii-Art-Stylize

This project is an enhancement of the Ascii Art Web project, aiming to make the website more appealing, interactive, and user-friendly through the use of CSS.

## 📄 Description

Ascii-Art-Stylize is a web-based ASCII art generator implemented in Go. It enhances the original Ascii Art Web project by incorporating modern CSS to make the interface more appealing, interactive, and user-friendly.

### ✨ Features

    ASCII art generation based on user input.
    Enhanced web-based interface for easy interaction.
    Error handling for invalid input and missing template files.
    ASCII art output displayed using HTML and CSS.
    Consistent, responsive, and interactive design.

## 🛠 How It Works

    The program takes user input text.
    Generates ASCII art based on the selected banner style.
    Displays the results on a web interface.
    Provides a modern, interactive, and user-friendly interface using CSS.

## 🔍 Implementation

    Go: Used for the backend logic, handling HTTP requests, and rendering templates.
    HTML: Structures the web interface.
    CSS: Styles the interface to make it appealing, responsive, and interactive.

### Algorithms

    The program uses Go's standard packages for HTTP handling, template rendering, and file operations.
    It validates user input for ASCII characters only and checks for the existence and integrity of ASCII art templates.
    The web interface allows users to input text and see the generated ASCII art.


## 🎨 HTML and CSS Usage
HTML

The HTML file is structured to create a simple and intuitive interface for user interaction.
CSS

The CSS file is used to style the HTML elements, making the interface visually appealing and responsive.
## 💾 How to Install

Ascii-Art-Stylize requires Go v1.22.2 to run.

    Clone the repository:    
    
    git clone https://github.com/krodgers/ascii-art-stylize.git

    cd ascii-art-stylize

## ▶️ Running the Program

To run the program, use the following command:

    go run .

    The terminal will prompt:

    Starting server at :8080

    To access the web interface, open your web browser and navigate to http://localhost:8080.

## 💡 Usage

#### Access the Web Interface:

 Open your web browser and navigate to the provided URL where the HTML file is hosted (i.e., http://localhost:8080).

#### Input Text:

You'll see a text area labeled "Enter Input Text Here...". Click inside this area and type the text you want to convert into ASCII art.

#### Select a Banner Style:

Choose a banner style from the dropdown menu labeled "Select Banner:". Options include:
    Standard: Basic ASCII art representation.
    Thinkertoy: A more elaborate ASCII art style.
    Shadow: ASCII art with a shadow effect.

#### Generate ASCII Art:

Click on the "Generate ASCII Art" button below the input form. This will process your input and display the ASCII art output below the form.

#### View ASCII Art Output:

Once generated, the ASCII art output will appear under the heading "ASCII Art Output:". It will be displayed inside a bordered box for easy viewing.

#### Reset Input:

To clear the input and start over, you can click on the "Reset" button located below the input form. This will refresh the page, clearing the text area and resetting the banner selection.

#### Interact with the Output:

You can scroll within the ASCII art box to view larger outputs and adjust the view as needed.

## ✍️ Authors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/rodgers-kaunda>
            <img src=https://learn.zone01kisumu.ke/git/avatars/aa19095145ab1ad43695e3cd3f7f3a5b?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=krodgers/>
            <br />
            <sub style="font-size:14px"><b>krodgers</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/swabri-musa-565350291?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3Buf0Ls4oWR2O2WLUMO5sIBg%3D%3D>
            <img src=https://learn.zone01kisumu.ke/git/avatars/bc7899a0aac2630a0a9b50bf330437a7?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=skanenje/>
            <br />
            <sub style="font-size:14px"><b>skanenje</b></sub>
        </a>
    </td>

</tr>
</table>

Feel free to contribute to this project by opening issues or submitting pull requests!!