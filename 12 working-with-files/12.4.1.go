  var b bytes.Buffer
	x := 0

	for {

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input := scanner.Text()
		if input == "exit" {
			break
		}

		x = x + 1
		a := strconv.Itoa(x)

		time := time.Now()
		t := time.String()
		t = t[:19]

		b.WriteString(a)
		b.WriteString(" ")
		b.WriteString(t)
		b.WriteString(" ")
		b.WriteString(input)
		b.WriteString("\n")
	
  }
  fileName := "log.txt"
if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
panic(err)
}
  }