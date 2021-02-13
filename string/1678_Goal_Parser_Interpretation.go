func interpret(command string) string {
    r := strings.ReplaceAll(command, "(al)", "al")
    r = strings.ReplaceAll(r, "()", "o")
    return r
}