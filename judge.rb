class JudgeSystem
  def initialize(user_name, language)
    @lang = language
    @user_name = user_name
  end

  def main
    file_open
  end

  def file_open
    File.open("./code/#{@user_name}.#{@lang}", "r") do |file|
      puts file.read
    end
  end
end
