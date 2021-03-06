require 'digest'
require 'aws-sdk'

BUCKET_NAME = 'dickeyxxx_dev'

TARGETS = [
  {os: 'darwin', arch: 'amd64'},
  {os: 'linux', arch: 'amd64'},
  {os: 'linux', arch: '386'},
  {os: 'windows', arch: 'amd64'},
  {os: 'windows', arch: '386'}
]

DIRTY = `git status 2> /dev/null | tail -n1`.chomp != 'nothing to commit, working directory clean'

task :run do
  build(nil, nil, './launcher')
  exec './launcher', *ARGV[1..-1]
end

task :build do
  FileUtils.mkdir_p 'dist'
  TARGETS.each do |target|
    path = "./dist/launcher_#{target[:os]}_#{target[:arch]}"
    puts "building #{path}..."
    build(target[:os], target[:arch], path)
  end
end

task :deploy => :build do
  raise 'dirty' if DIRTY
  puts "deploying..."
  bucket = get_s3_bucket
  TARGETS.each do |target|
    filename = "launcher_#{target[:os]}_#{target[:arch]}"
    local_path = "./dist/#{filename}"
    remote_path = "launcher/#{filename}"
    remote_url = "#{BUCKET_NAME}.s3.amazonaws.com/#{remote_path}"
    puts "uploading #{local_path} to #{remote_url}"
    upload_file(bucket, local_path, remote_path)
  end
end

def build(os, arch, path)
  system("GOOS=#{os} GOARCH=#{arch} go build -o #{path}")
end

def write_digest(path)
  digest = Digest::SHA1.file(path).hexdigest
  File.open(path + '.sha1', 'w') { |f| f.write(digest) }
end

def get_s3_bucket
  s3 = AWS::S3.new
  s3.buckets[BUCKET_NAME]
end

def upload_file(bucket, local, remote)
  obj = bucket.objects[remote]
  obj.write(Pathname.new(local))
  obj.acl = :public_read
end
