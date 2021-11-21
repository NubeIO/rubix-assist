class Upload {
  final String? fromPath;
  final String? toPath;
  final String? unzipPath;
  final bool? unzip;
  final bool? clearDir;
  final List<String>? zips;

  Upload(
      {this.fromPath,
      this.toPath,
      this.unzipPath,
      this.unzip,
      this.clearDir,
      this.zips});

  factory Upload.fromJson(Map<String, dynamic> json) {
    return Upload(
      fromPath: json["from_path"],
      toPath: json["to_path"],
      unzipPath: json["unzip_path"],
      unzip: json["unzip"],
      clearDir: json["clear_dir"],
    );
  }

  @override
  String toString() {
    // TODO: implement toString
    return 'Host{id: $fromPath, name: $toPath}';
  }
}
