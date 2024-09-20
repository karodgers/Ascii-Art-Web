# Ascii Art Web
## Description
The ASCII Art Web (generator) is a web application built using the Go programming language. It allows users to convert their input text into ASCII art, leveraging predefined templates for various styles. This program ensures a smooth and efficient user experience by utilizing Go's standard packages for HTTP handling, template rendering, and file operations.
### Features
- HTTP Handling: The application uses Go's net/http package to manage server operations, handling incoming requests and responses efficiently.
- Template Rendering: The html/template package is employed to render HTML templates. This ensures dynamic content generation and a responsive web interface.
- File Operations: Go's os package is utilized for reading and validating ASCII art templates, ensuring the templates are available and intact.
- ASCII art output displayed using HTML and CSS.

### Authors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/swabri-musa-565350291?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3Buf0Ls4oWR2O2WLUMO5sIBg%3D%3D>
            <img src=https://learn.zone01kisumu.ke/git/avatars/bc7899a0aac2630a0a9b50bf330437a7?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=skanenje/>
            <br />
            <sub style="font-size:14px"><b>skanenje</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/rodgers-kaunda>
            <img src=https://learn.zone01kisumu.ke/git/avatars/aa19095145ab1ad43695e3cd3f7f3a5b?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=krodgers/>
            <br />
            <sub style="font-size:14px"><b>krodgers</b></sub>
        </a>
    </td>
</tr>
</table>

## Usage
### How to Install
Ascii-art requires [go](https://go.dev/dl/)  v 1.22.2 to run

Once go is installed  clone into this repository to do this 

- Open terminal and run
``` sh
git clone https://github.com/karodgers/ascii-art-web
cd ascii-art-web
```
Once you are in the directory,

To run the program, use the following command, on the terminal command line:
```bash
go run main.go
```
The terminal will prompt 
```
Starting server at :8080
to access visit http://localhost:8080
```
Use the link provided to the access the web GUI.

1. **Access the Web Interface:**

    Open your web browser and navigate to the provided URL where the HTML file is hosted (i.e, http://localhost:8080/).

2. **Input Text:**

    You'll see a text area labeled "Enter Input Text Here:". Click inside this area and type the text you want to convert into ASCII art.

3. **Select a Banner Style:**

    Choose a banner style from the dropdown menu labeled "Select Banner:". Options include:
        Standard: Basic ASCII art representation.
        Thinkertoy: A more elaborate ASCII art style.
        Shadow: ASCII art with a shadow effect.

4. **Generate ASCII Art:**

    Click on the "Generate ASCII Art" button below the input form. This will process your input and display the ASCII art output below the form.

5. **View ASCII Art Output:**

    Once generated, the ASCII art output will appear under the heading "ASCII Art Output:". It will be displayed inside a bordered box for easy viewing.

7. **Reset Input:**

    To clear the input and start over, you can click on the "Reset" button located below the input form. This will refresh the page, clearing the text area and resetting the banner selection.

8. **Interact with the Output:**

    You can scroll within the ASCII art box to view larger outputs and adjust the view as needed.
## Implementation: Algorithms
- The program uses Go's standard packages for HTTP handling, template rendering, and file operations.
- It validates user input for ASCII characters only and checks for the existence and integrity of ASCII art templates.
- The web interface allows users to input text, and see the generated ASCII art.

#### User Input Validation

- To maintain the integrity and security of the application, user inputs are validated to contain only ASCII characters. This prevents potential issues related to non-ASCII input and ensures consistent output.

#### Template Integrity Check

- The application checks for the existence and integrity of ASCII art templates before processing any input. This validation step guarantees that the templates are correctly formatted and available, preventing errors during the ASCII art generation process.

#### Web Interface

- The user-friendly web interface allows users to easily interact with the application. Users can input text into a form and submit it to see the generated ASCII art. The interface is designed to be intuitive, providing a seamless experience from input to output.


