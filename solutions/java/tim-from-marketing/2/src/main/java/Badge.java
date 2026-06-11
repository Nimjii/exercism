class Badge {
    public String print(Integer id, String name, String department) {
        return (id != null ? "[" + id.toString() + "] - " : "")
            + name
            + (department != null ? " - " + department.toUpperCase() : " - OWNER");
    }
}
