from setuptools import setup, find_packages

install_requires=[
    "oauthlib==3.2.2",
    "requests==2.30.0",
    "requests-oauthlib==1.3.1",
    "urllib3==2.0.2"
    ]
setup(
    name='druvareportsdk',
    version="1.0.1",
    description='Python SDK for Druva Reports.',
    author='Team DCP',
    author_email='support@druva.com',
    url='https://github.com/druvainc/Platform/reporting-api/pysdk',
    license='Druva',
    python_requires='>=3.0',
    packages=find_packages(),
    include_package_data=True,
    install_requires=install_requires
)